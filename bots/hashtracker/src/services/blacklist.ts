import { gql, request } from 'graphql-request';

const url = 'https://api.studio.thegraph.com/query/91524/blacklist-index/version/latest'

export async function getBlacklist(
    address:string
): Promise<any> {
    const query = gql`
      query GetReportsByAddress($address: String!) {
        reportCreateds(where: { reportedAddress: $address }) {
          id
          reportedAddress
          count
          category
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
    if (report.reportCreateds.length) {
        return {
            isBlacklisted :true,
            report: report
        }
    }

     return {
        isBlacklisted :false,
    }
}
