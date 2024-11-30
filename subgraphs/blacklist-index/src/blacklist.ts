import { Blacklisted as BlacklistedEvent } from "../generated/Blacklist/Blacklist"
import { Blacklisted } from "../generated/schema"

export function handleBlacklisted(event: BlacklistedEvent): void {
  let entity = new Blacklisted(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.reportedAddress = event.params.reportedAddress
  entity.count = event.params.count
  entity.category = event.params.category
  entity.date = event.params.date
  entity.comments = event.params.comments
  entity.source = event.params.source

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
