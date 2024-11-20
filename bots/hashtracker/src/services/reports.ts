
export async function createReport(
    address:string
): Promise<any> {
    console.log("--> querying transactions: ", address);
    const baseUrl = 'https://hashtracker-em3r7.ondigitalocean.app';
    const url = `${baseUrl}/reports/${address}`;
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
            },
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const message = await response.json();
        const data =  JSON.stringify(message, null, 2);
        return JSON.parse(data);
    } catch (error) {
        throw new Error(`Failed to fetch wallet transactions: ${error}`);
    }
}
