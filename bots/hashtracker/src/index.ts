import {
    run,
    XMTPContext,
    agentReply,
    replaceVariables, Skill, Agent,
} from "@xmtp/message-kit";
import { transactionSkills} from "./skills.js";
import { systemPrompt } from "./prompt.js";


export const agent: Agent = {
    name: "Transactional tracking Bot",
    description: "Transactional tracking Bot",
    tag: "@hashi",
    skills: [
        ...transactionSkills,
    ],
};

run(
    async (context:XMTPContext) => {
        const {
            message: { sender },
            runConfig,
        } = context;

        let prompt = await replaceVariables(systemPrompt, sender.address, agent);
        await agentReply(context, prompt);
    },
    { agent },
);
