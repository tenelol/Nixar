// apps/portfolio/script.js
document.addEventListener("DOMContentLoaded", async () => {
  const container = document.getElementById("projects");
  if (!container) return; // projectsセクションがないページでは何もしない

  container.textContent = "Loading...";

  try {
    const res = await fetch("/api/projects");
    if (!res.ok) {
      throw new Error("failed to fetch");
    }

    const projects = await res.json();

    if (!Array.isArray(projects) || projects.length === 0) {
      container.textContent = "No projects yet.";
      return;
    }

	const html = projects
  	 .map(
	   (p) => `
      		<article class="project-card">
        	<h2>${p.name}</h2>
       	 	<p>${p.description}</p>
        	${
          	p.url
            	? `<p><a href="${p.url}" target="_blank" rel="noopener noreferrer">View more</a></p>`
            : ""
        }
      </article>
    `
  )
  .join("");


    container.innerHTML = html;
  } catch (err) {
    console.error(err);
    container.textContent = "Failed to load projects.";
  }
});

// 例：loadProjects の中

projects.forEach(p => {
    const card = document.createElement('article');
    card.className = 'project-card';

    card.innerHTML = `
        <h3>${p.name}</h3>
        <p>${p.description}</p>
        ${p.url ? `<a class="card-link" href="${p.url}" target="_blank" rel="noopener">View</a>` : ''}
    `;

    container.appendChild(card);
});

