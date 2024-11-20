import {XMTPContext, SkillResponse} from "@xmtp/message-kit";
import {fetchTransactions} from "../services/transactions.js";


export async function chatHandler(
    context: XMTPContext,
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
        const url = `https://hashtracker.vercel.app/?hash=${transactions.mermaid}`
        return { code: 200, message: `The following diagram illustrates the ${address} address behaviour. ${url}` };
    }

    if (skill == "report") {
        const { address } = params;
        return { code: 200, message: `${address} reported...` };
    }
}
