/**
* This code was generated by v0 by Vercel.
* @see https://v0.dev/t/ZFxL7Nxs5ve
* Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
*/
"use client";

import React, { useState } from "react"
import { SelectValue, SelectTrigger, SelectItem, SelectGroup, SelectContent, Select } from "@/components/ui/select"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
const axios = require('axios');

export default function Home() {
  const [db,setDb] = useState<string>("");
  const [promptText, setPromptText] = useState<string>("");
  const [responseCode, setResponseCode] = useState<string>("");
  async function sendPrompt(e:any) {
    e.preventDefault();
    const response = await fetch('/api/v1/dbquery/firebase', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        database: db,
        prompt: promptText,
      }),
    });
    const data = await response.json();
    setResponseCode(data.result);
  }
  return (
    <section className="grid grid-cols-2 gap-20 p-40 bg-blue-200 h-screen">
      <div className="space-y-4">
        <h2 className="text-3xl font-bold text-black">DBQuery.ai</h2>
        <p className="text-black text-2xl">Demo 💪🏻</p>
        <form className="space-y-4">
          <Select onValueChange={(e)=>setDb(e)}>
            <SelectTrigger className="w-full">
              <SelectValue placeholder="Choose your DB" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="Firebase">Firebase</SelectItem>
                <SelectItem value="Supabase">Supabase</SelectItem>
                <SelectItem value="Postgresql">Postgresql</SelectItem>
                <SelectItem value="MongoDB">MongoDB</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <Textarea placeholder="Your message" className="text-lg" onChange={(e:any)=>setPromptText(e.target.value)}/>
          <Button className="w-full" type="submit" onClick={sendPrompt}>
            Submit
          </Button>
        </form>
      </div>
      <div className="bg-gray-100 dark:bg-gray-800 rounded-xl p-8 flex justify-center">
        <div className="max-w-[500px] text-center space-y-4">
          <h3 className="text-2xl font-bold">Results:</h3>
          {responseCode && <p className="text-lg">{responseCode}</p>}
        </div>
      </div>
    </section>
  )
}

function ArrowRightIcon(props:any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M5 12h14" />
      <path d="m12 5 7 7-7 7" />
    </svg>
  )
}
