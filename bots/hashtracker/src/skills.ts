import {chatHandler} from "./handler/chat.js";

import type { SkillGroup } from "@xmtp/message-kit";

export const skills: SkillGroup[] = [
    {
        name: "Ens Domain Bot",
        tag: "@ens",
        description: "Register ENS domains.",
        skills: [
            {
                skill: "/report [address]",
                description: "Report an address as dangerous.",
                handler: chatHandler,
                examples: ["/report 0x1234567890123456789012345678901234567890"],
                params: {
                    address: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/transactions [address]",
                description: "Show the transactions of a specific address.",
                handler: chatHandler,
                examples: ["/transactions 0x1234567890123456789012345678901234567890"],
                params: {
                    address: {
                        type: "string",
                    },
                },
            },
        ],
    },
];
