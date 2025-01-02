import {
    run,
    Agent,
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
    systemPrompt: systemPrompt,
};

run( agent );
