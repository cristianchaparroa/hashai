import { HandlerContext, SkillResponse } from "@xmtp/message-kit";
import { getUserInfo, clearInfoCache, isOnXMTP } from "@xmtp/message-kit";
import { isAddress } from "viem";
import { clearMemory } from "@xmtp/message-kit";


export const ensUrl = "https://app.ens.domains/";


export async function handleEns(
    context: HandlerContext,
): Promise<SkillResponse | undefined> {
    const {
        message: {
            sender,
            content: { skill, params },
        },
    } = context;
    console.log(skill, params);
    if (skill == "info") {
        const { domain } = params;

        const data = await getUserInfo(domain);
        if (!data?.ensDomain) {
            return {
                code: 404,
                message: "Domain not found.",
            };
        }

        const formattedData = {
            Address: data?.address,
            "Avatar URL": data?.ensInfo?.avatar,
            Description: data?.ensInfo?.description,
            ENS: data?.ensDomain,
            "Primary ENS": data?.ensInfo?.ens_primary,
            GitHub: data?.ensInfo?.github,
            Resolver: data?.ensInfo?.resolverAddress,
            Twitter: data?.ensInfo?.twitter,
            URL: `${ensUrl}${domain}`,
        };

        let message = "Domain information:\n\n";
        for (const [key, value] of Object.entries(formattedData)) {
            if (value) {
                message += `${key}: ${value}\n`;
            }
        }
        return { code: 200, message };

    } else {
        return { code: 400, message: "Skill not found." };
    }
}


export async function clear() {
    clearMemory();
    clearInfoCache();
}
