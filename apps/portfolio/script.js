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
function getFavicon(url) {
  try {
    const domain = new URL(url).hostname;
    // SVG favicon API に変更
    return `https://icons.duckduckgo.com/ip3/${domain}.ico`;
  } catch {
    // URL おかしいとき用の保険
    return "/static/default-icon.png";
  }
}




const html = projects
  .map(
    (p) => `
      <a class="project-card"
         href="${p.url}"
         target="_blank"
         rel="noopener noreferrer">
        <div class="project-card-inner">
          <div class="project-icon">
            <img src="${getFavicon(p.url)}" alt="${p.name} icon" />
          </div>
          <div class="project-content">
            <h2>${p.name}</h2>
            <p>${p.description}</p>
          </div>
        </div>
      </a>
    `
  )
  .join("");

container.innerHTML = html;



    container.innerHTML = html;
  } catch (err) {
    console.error(err);
    container.textContent = "Failed to load projects.";
  }
});


