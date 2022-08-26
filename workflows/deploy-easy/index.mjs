// Build todoapp, the EASY way
import { gql, Engine } from "@dagger.io/dagger";

new Engine({
	ConfigPath: process.env.CLOAK_CONFIG
}).run(async (client) => {
	// 1. Load app source code from working directory
	const source = await client.request(gql`
	{
	  host {
	    workdir {
	      read {
	        id
	      }
	    }
	  }
	}
  `).then((result) => result.host.workdir.read)

	// 2. Install yarn in a container
	const sourceAfterBuild = await client.request(gql`
	{
	  yarn {
		script(source: "${source.id}", runArgs: ["react-scripts", "build"]) {
			id
		}
	  }
	}
	`).then((result) => result.yarn.script)

	// 3. set netlify token as secret (read from the env)
	const netlifyToken = await client.request(gql`
	{
		core {
			addSecret(plaintext: "${process.env.NETLIFY_AUTH_TOKEN}")
		}
	}
	`).then((result) => result.core.addSecret)

	// 4. deploy to Netlify
	const result = await client.request(gql`
	{
		netlify {
			deploy(contents: "${sourceAfterBuild.id}", subdir: "build", siteName: "sam-cloak-test-demo", token: "${netlifyToken}") {
				url
				deployURL
			}
		}
	}
	`)

	console.log("Netlify deploy URL: " + result.netlify.deploy.url)
});
