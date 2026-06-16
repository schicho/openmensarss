(async () => {
    const ghResp = await fetch("https://api.github.com/repos/schicho/openmensarss/contents/rss");
    const ghData = await ghResp.json();
    const idList = ghData
        .filter((file) => file.name.endsWith(".xml"))
        .map((file) => file.name.split(".", 1)[0])
        .join(",");
    const omResp = await fetch(`https://openmensa.org/api/v2/canteens?ids=${idList}`);
    const canteenList = await omResp.json();

    const ul = document.createElement("ul");
    for (let canteen of canteenList) {
        const filename = `${canteen.id}.xml`;
        const li = document.createElement("li");
        li.innerHTML = `<p><a href="${filename}">${filename}</a> ${canteen.name}</p>`;
        ul.appendChild(li)
    }
    document.getElementById("filelist").innerHTML = "";
    document.getElementById("filelist").appendChild(ul);
})();
