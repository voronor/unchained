import { Blockbook } from '@shapeshiftoss/blockbook'
import { Message, Worker, SyncTx } from '@shapeshiftoss/common-ingester'
import { logger } from '../logger'
import { parseTx } from '../parseTx'
import { SequencedBTCParseTx } from '../types'

const INDEXER_URL = process.env.INDEXER_URL
const INDEXER_WS_URL = process.env.INDEXER_WS_URL

if (!INDEXER_URL) throw new Error('INDEXER_URL env var not set')
if (!INDEXER_WS_URL) throw new Error('INDEXER_WS_URL env var not set')

const blockbook = new Blockbook({ httpURL: INDEXER_URL, wsURL: INDEXER_WS_URL })

const msgLogger = logger.child({ namespace: ['workers', 'address'], fn: 'onMessage' })
const onMessage = (worker: Worker) => async (message: Message) => {
  const { address, txid, client_id, sequence, total }: SyncTx = message.getContent()
  const retryKey = `${client_id}:${address}:${txid}`

  try {
    const tx = await blockbook.getTransaction(txid)
    msgLogger.trace({ blockHash: tx.blockHash, blockHeight: tx.blockHeight, txid: tx.txid }, 'Transaction')

    const pTx = await parseTx(tx, address)

    worker.sendMessage(new Message({ ...pTx, sequence, total } as SequencedBTCParseTx), client_id)
    worker.ackMessage(message, retryKey)

    msgLogger.debug({ address, txid, client_id }, 'Transaction published')
  } catch (err) {
    msgLogger.error(err, 'Error processing address')
    worker.retryMessage(message, retryKey)
  }
}

const main = async () => {
  const worker = await Worker.init({
    queueName: 'queue.txid.address',
    exchangeName: 'exchange.tx.client',
  })

  worker.queue?.prefetch(100)
  worker.queue?.activateConsumer(onMessage(worker), { noAck: false })
}

main().catch((err) => {
  logger.error(err)
  process.exit(1)
})
