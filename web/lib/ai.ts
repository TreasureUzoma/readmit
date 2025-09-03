import { GoogleGenAI } from "@google/genai"

export const systemPrompts = {
  readme: `You are a helpful assistant that generates a professional README file based on the provided project content.`,
  dockerfile: `You are a helpful assistant that generates a Dockerfile tailored to the provided project content.`,
  commits: `You are a helpful assistant that suggests meaningful Git commit messages for the provided project content.`,
  contributionmd: `You are a helpful assistant that generates a CONTRIBUTING.md file based on the provided project content.`,
}

const ai = new GoogleGenAI({
  apiKey: process.env.GOOGLE_GENAI_API_KEY!,
})

interface GenAIOptions {
  mode: keyof typeof systemPrompts
  input: string
}

export const genai = async ({ mode, input }: GenAIOptions): Promise<string> => {
  const systemInstruction = systemPrompts[mode]

  if (!systemInstruction) {
    throw new Error(`Unsupported mode: ${mode}`)
  }

  const isUrl = input.startsWith("http://") || input.startsWith("https://")

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: systemInstruction + `Here is the codebase url: ${input}`,
    config: {
      tools: [{ urlContext: {} }],
    },
  })

  console.log("AI Response:", response.text)
  return response.text!
}
