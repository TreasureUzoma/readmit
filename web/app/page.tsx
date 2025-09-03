import Link from "next/link"

import { PageRoutes } from "@/lib/pageroutes"
import { buttonVariants } from "@/components/ui/button"

export default function Home() {
  return (
    <div>
      <svg
        className="fixed inset-0 w-full h-full text-muted -z-10 opacity-[0.04] dark:opacity-[0.07]"
        xmlns="http://www.w3.org/2000/svg"
      >
        <defs>
          <pattern
            id="dotPattern"
            width="40"
            height="40"
            patternUnits="userSpaceOnUse"
          >
            <circle cx="1" cy="1" r="1" fill="currentColor" />
          </pattern>
        </defs>
        <rect width="100%" height="100%" fill="url(#dotPattern)" />
      </svg>

      {/* Hero Section */}
      <section className="min-h-[90vh] flex flex-col justify-center items-center text-center px-4 py-24 md:py-30">
        <div className="max-w-4xl mx-auto">
          <h1 className="text-4xl md:text-6xl font-bold mb-6 tracking-tight">
            AI-powered documentation generator
          </h1>
          <p className="text-xl text-muted-foreground mb-6 max-w-3xl mx-auto leading-relaxed">
            Generate READMEs, contribution guides, commit messages, Dockerfiles,
            and more—automatically.
          </p>

          <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mb-16">
            <Link
              href="/docs/introduction"
              className={`${buttonVariants({
                size: "lg",
              })} w-[70%] font-medium`}
            >
              Read the Documentation
            </Link>
          </div>

          {/* Code Example */}
          <div className="bg-muted text-muted-foreground p-6 rounded-xl font-mono text-left max-w-2xl mx-auto mb-16 shadow-sm">
            <div className="mb-2"># Install readmit (macOS / Linux)</div>
            <div className="text-green-500">
              $ curl -fsSL https://readmit.vercel.app/install | bash
            </div>

            <div className="mt-4"># Install readmit (Windows PowerShell)</div>
            <div className="text-green-500">
              PS&gt; irm https://readmit.vercel.app/install.ps1 | iex
            </div>

            <div className="mt-6"># Generate documentation</div>
            <div className="text-green-500">$ readmit generate readme</div>
            <div className="mt-2">✓ Analyzing codebase...</div>
            <div>✓ Generated README.md</div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 px-4 border-t border-border">
        <div className="max-w-6xl mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold text-center mb-16">
            What readmit generates
          </h2>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {[
              {
                title: "README.md",
                description:
                  "Comprehensive project documentation with installation, usage, and API references",
              },
              {
                title: "CONTRIBUTING.md",
                description:
                  "Contributor guidelines based on your project structure and coding patterns",
              },
              {
                title: "Commit Messages",
                description:
                  "Intelligent commit message suggestions following conventional commit standards",
              },
              {
                title: "Dockerfiles",
                description:
                  "Optimized Docker configurations tailored to your tech stack",
              },
              {
                title: "API Documentation",
                description:
                  "Auto-generated API docs from your route handlers and endpoints",
              },
              {
                title: "License",
                description:
                  "License information and guidelines for your project",
                comingSoon: true,
              },
            ].map((feature, index) => (
              <div
                key={index}
                className="rounded-xl border border-border p-6 bg-card hover:bg-accent hover:text-accent-foreground transition-colors duration-200 shadow-sm"
              >
                <span className="text-sm text-muted-foreground">
                  {feature.comingSoon && "Coming Soon"}
                </span>
                <h3 className="text-xl font-bold mb-3">{feature.title}</h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* How it works */}
      <section className="py-20 px-4 bg-muted">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="text-4xl md:text-5xl font-bold mb-16">How it works</h2>

          <div className="grid md:grid-cols-3 gap-12">
            {[
              {
                step: "01",
                title: "Scan",
                desc: "Reads your entire codebase, understanding structure, dependencies, and patterns",
              },
              {
                step: "02",
                title: "Analyze",
                desc: "AI processes your code to understand functionality, architecture, and best practices",
              },
              {
                step: "03",
                title: "Generate",
                desc: "Creates comprehensive, accurate documentation tailored to your specific project",
              },
            ].map((s, i) => (
              <div key={i} className="space-y-4">
                <div className="text-6xl font-bold text-muted-foreground">
                  {s.step}
                </div>
                <h3 className="text-2xl font-bold">{s.title}</h3>
                <p className="text-muted-foreground leading-relaxed">
                  {s.desc}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 px-4 border-t border-border">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Stop writing docs manually
          </h2>
          <p className="text-xl text-muted-foreground mb-12 leading-relaxed">
            Let readmit understand your code and generate professional
            documentation in seconds.
          </p>

          <Link
            href={`/docs${PageRoutes[0].href}`}
            className={buttonVariants({
              size: "lg",
            })}
          >
            Start Generating Docs
          </Link>
        </div>
      </section>
    </div>
  )
}
