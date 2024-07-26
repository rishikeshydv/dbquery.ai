// await axios.post('http://54.83.89.143:9002/api/v1/dbquery/firebase', 
// {
//   database: db,
//   prompt: promptText
// }
import { NextRequest, NextResponse } from 'next/server';
import axios from 'axios';
export async function POST(req: NextRequest, res: NextResponse) {
    try {
      const reqData = await req.json();
      const response = await axios.post('http://54.83.89.143:9002/api/v1/dbquery/firebase', {
        database: reqData.database,
        prompt: reqData.prompt
      });
      return NextResponse.json({ result: response.data.result });
    } catch (error) {
      console.error('Error during API call:', error);
      return NextResponse.json({ error: 'An error occurred while processing your request.' }, { status: 500 });
    }
  }