import { ReportCreated as ReportCreatedEvent } from "../generated/HashReporter/HashReporter"
import { ReportCreated } from "../generated/schema"

export function handleReportCreated(event: ReportCreatedEvent): void {
  let entity = new ReportCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.reportedAddress = event.params.reportedAddress
  entity.count = event.params.count
  entity.category = event.params.category

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
