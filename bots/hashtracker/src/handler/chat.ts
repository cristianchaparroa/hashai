import {HandlerContext, SkillResponse} from "@xmtp/message-kit";


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
        return { code: 200, message: `Diagram for the address ${address}.` };
    }

    if (skill == "report") {
        const { address } = params;
        return { code: 200, message: `${address} reported...` };
    }
}
