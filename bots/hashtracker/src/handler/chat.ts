import {HandlerContext, SkillResponse} from "@xmtp/message-kit";
import {fetchTransactions} from "../services/transactions.js";


export async function chatHandler(
    context: HandlerContext,
): Promise<SkillResponse | undefined> {
    const {
        message: {
            sender,
            content: {skill, params},
        },
    } = context;
    console.log(skill, params);

    if (skill == "transactions") {
        const { address } = params;
        const transactions = await fetchTransactions(address);
        return { code: 200, message: `The following diagram illustrates the behaviour of ${address} address. ${transactions.mermaid}` };
    }

    if (skill == "report") {
        const { address } = params;
        return { code: 200, message: `${address} reported...` };
    }
}
