import {
    run,
    XMTPContext,
    agentReply,
    replaceVariables, SkillGroup,
} from "@xmtp/message-kit";
import { transactionSkills} from "./skills.js";
import { systemPrompt } from "./prompt.js";


export const skills = [
    {
        name: "Transactional tracking Bot",
        tag: "@hashi",
        description: "Transactional tracking Bot",
        skills: [
            ...transactionSkills,
        ],
    },
];

run(
    async (context:XMTPContext) => {
        const {
            message: { sender },
            runConfig,
        } = context;

        let prompt = await replaceVariables(
            systemPrompt,
            sender.address,
            runConfig?.skills,
            "@hashi",
        );
        await agentReply(context, prompt);
    },
    { skills },
);
