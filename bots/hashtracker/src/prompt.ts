import { skills } from "./skills.js";
import { defaultPromptTemplate } from "@xmtp/message-kit";


export async function agent_prompt(senderAddress: string) {
  let fineTuning = `
## Example responses:
 
1. Check the transactions of the wallet address
  Hey {PREFERRED_NAME}! the following are the transactions generated for the address {MERMAID_B64}
 
2. Hey I want to report the next address because it stolen my founds 
  Hello  {PREFERRED_NAME} the report for the {ADDRESS} address has been completed.
`;

  return defaultPromptTemplate(fineTuning, senderAddress, skills, "@ens");
}
