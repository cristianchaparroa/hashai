import { handleEns } from "./handler/ens.js";
import {chatHandler} from "./handler/chat.js";

import type { SkillGroup } from "@xmtp/message-kit";

export const skills: SkillGroup[] = [
    {
        name: "Ens Domain Bot",
        tag: "@ens",
        description: "Register ENS domains.",
        skills: [
            {
                skill: "/register [domain]",
                handler: handleEns,
                description:
                    "Register a new ENS domain. Returns a URL to complete the registration process.",
                examples: ["/register vitalik.eth"],
                params: {
                    domain: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/info [domain]",
                handler: handleEns,
                description:
                    "Get detailed information about an ENS domain including owner, expiry date, and resolver.",
                examples: ["/info nick.eth"],
                params: {
                    domain: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/renew [domain]",
                handler: handleEns,
                description:
                    "Extend the registration period of your ENS domain. Returns a URL to complete the renewal.",
                examples: ["/renew fabri.base.eth"],
                params: {
                    domain: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/check [domain]",
                handler: handleEns,
                examples: ["/check vitalik.eth", "/check fabri.base.eth"],
                description: "Check if a domain is available.",
                params: {
                    domain: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/cool [domain]",
                examples: ["/cool vitalik.eth"],
                handler: handleEns,
                description: "Get cool alternatives for a .eth domain.",
                params: {
                    domain: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/reset",
                examples: ["/reset"],
                handler: handleEns,
                description: "Reset the conversation.",
                params: {},
            },
            {
                skill: "/tip [address]",
                description: "Show a URL for tipping a domain owner.",

                handler: handleEns,
                examples: ["/tip 0x1234567890123456789012345678901234567890"],
                params: {
                    address: {
                        type: "string",
                    },
                },
            },
            {
                skill: "/report [address]",
                description: "Report an address as dangerous.",
                handler: chatHandler,
                examples: ["/tip 0x1234567890123456789012345678901234567890"],
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