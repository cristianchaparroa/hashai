import {reportHandler, transactionsHandler} from "./handler/handler.js";
import { SkillGroup, SkillParamConfig } from "@xmtp/message-kit";

export const transactionSkills:SkillGroup["skills"] =  [
    {
        skill: "/report [addres]",
        handler: reportHandler,
        examples:  [
            "/report 0x1234567890123456789012345678901234567890",
            "/report vitalik.eth",
        ],
        description: "Report an address as dangerous.",
        params: {
            address: {
                type: "string",
            } as SkillParamConfig,
        },
    },

    {
        skill: "/transactions [address]",
        handler: transactionsHandler,
        examples: [
            "/transactions vitalik.eth",
            "/transactions fabri.base.eth",
            "transactions0x1234567890123456789012345678901234567890"
        ],
        description: "Show the transactions of a specific address",
        params: {
            address: {
                type: "string",
            } as SkillParamConfig,
        },
    },
];
