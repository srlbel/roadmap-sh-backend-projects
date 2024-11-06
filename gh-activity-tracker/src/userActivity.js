import figlet from "figlet";

export const fetchUserActivity = (username) => {
  fetch(`https://api.github.com/users/${username}/events`)
    .then((response) => {
      if (!response.ok) {
        if (response.status == 404) {
          throw Error("No user found")
        } else {
          throw Error("Unable to fetch resources.")
        }
      }

      console.log(figlet.textSync(username))
      return response.json()
    })
    .then((data) => formatInfo(data))
    .catch((err) => {
      console.error(err)
    })
}

const formatInfo = (data) => {
  data.forEach(element => {
    switch (element.type) {
      case "PushEvent":
        console.log(` - Pushed ${element.payload.commits.length} commits to ${element.repo.name}. \n`)
        break;

      case "CreateEvent":
        console.log(` - Created a new repo at ${element.repo.name}. \n`)
        break;

      case "DeleteEvent":
        console.log(` - Deleted repo ${element.repo.name}. \n`)
        break;

      case "WatchEvent":
        console.log(` - Turn watched repo ${element.repo.name}. \n`)
        break;

      case "PullRequestEvent":
        console.log(` - Openend a pull request at ${element.repo.name}. \n`)
        break;

      case "ReleaseEvent":
        console.log(` - Made a release at ${element.repo.name}. \n`)
        break;

      case "IssuesEvent":
        console.log(` - Opened an issue at ${element.repo.name}. \n`)
        break;

      case "ForkEvent":
        console.log(` - Forked a repo from ${element.repo.name} to ${element.payload.forkee.full_name}. \n`)
        break;
    }
  });
}