import { gql, request } from 'graphql-request';

const url = 'https://api.studio.thegraph.com/query/91524/blacklist-index/v0.0.10-3'

export async function getBlacklist(
    address:string
): Promise<any> {
    const query = gql`
      query GetReportsByAddress($address: String!) {
        blacklisteds(where: { reportedAddress: $address }) {
          id
          reportedAddress
          count
          category
          comments
        }
      }
    `;

    const variables = {
        address: address.toLowerCase()
    };

    let result = await request(url, query, variables);
    const data =  JSON.stringify(result, null, 2);
    let report = JSON.parse(data);
    console.log(report);
    if (report.blacklisteds.length) {
        return {
            isBlacklisted :true,
            report: report
        }
    }

     return {
        isBlacklisted :false,
    }
}
