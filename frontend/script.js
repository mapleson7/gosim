createPedigree();
let currentHorseID = 45;

getHorse(currentHorseID).then((horse) => {
  updateHorseData(horse);
});

document.getElementById("left-arr").addEventListener("click", () => {
  currentHorseID -= 1;
  getHorse(currentHorseID).then((horse) => updateHorseData(horse));
});
document.getElementById("right-arr").addEventListener("click", () => {
  currentHorseID += 1;
  getHorse(currentHorseID).then((horse) => updateHorseData(horse));
});

function updateHorseData(horse) {
  console.log(horse);
  removeGraphicalStatBars();
  createGraphicalStatBars([
    horse.Speed,
    horse.Acceleration,
    horse.Stamina,
    horse.Consistency,
    horse.Versatility,
  ]);
  removeTalentBoxes();
  applyTalentColourSchemeToStatBars([
    horse.SpeedTalent,
    horse.AccelerationTalent,
    horse.StaminaTalent,
    horse.ConsistencyTalent,
    horse.VersatilityTalent,
  ]);
  document.getElementById("horse-name").innerText = horse.Name;
  document.getElementById("sex").innerText = horse.Sex.toUpperCase();
}

function removeTalentBoxes() {
  const talentBoxes = document.querySelectorAll(".talent-box");
  for (let x = 0; x < talentBoxes.length; x++) {
    talentBoxes[x].remove();
  }
}

function applyTalentColourSchemeToStatBars(talentRarity) {
  //Depending on the Talent Rarity apply the appropriate color
  const statBoxes = document.querySelectorAll(".stat-box");

  for (let x = 0; x < statBoxes.length; x++) {
    const talentBox = document.createElement("div");
    let talentColor;
    let text;
    switch (talentRarity[x]) {
      case "none":
        text = "+0%";
        talentBox.style.visibility = "hidden";
        talentColor = "20, 20, 20";
        break;
      case "common":
        talentColor = "34, 139, 34";
        text = "+2.5%";
        break;
      case "rare":
        talentColor = "30, 144, 255";
        text = "+5%";
        break;
      case "epic":
        talentColor = "106, 13, 173";
        text = "+7.5%";
        break;
      case "legendary":
        talentColor = "255, 215, 30";
        text = "+10%";
        break;
    }
    talentBox.style.borderLeft = `8px solid rgba(${talentColor}, 0.8)`;
    talentBox.style.borderRight = `8px solid rgba(${talentColor}, 0.8)`;
    talentBox.style.backgroundColor = `rgb(${talentColor})`;
    talentBox.style.borderRadius = "20px";
    talentBox.innerText = text;
    talentBox.style.fontSize = "1.2rem";
    talentBox.classList.add("talent-box");

    statBoxes[
      x
    ].lastElementChild.style.borderRight = `4 px solid rgba(${talentColor}, 0.8)`;
    statBoxes[x].style.borderBottom = `4px solid rgba(${talentColor}, 0.8)`;
    statBoxes[x].style.borderRight = `2px solid rgba(${talentColor}, 0.8)`;
    statBoxes[x].style.backgroundColor = `rgba(${talentColor}, 0.75)`;
    statBoxes[x].insertBefore(talentBox, statBoxes[x].firstChild);
  }
}

async function getHorse(id) {
  const url = `http://localhost:3000/horses?id=${id}`;
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const responseJSON = await response.json();
    return responseJSON;
  } catch (err) {
    console.error(err.message);
  }
}

function removeGraphicalStatBars() {
  const statGraphicalSections = document.querySelectorAll(".stat-graphical");

  for (
    let sectionIndex = 0;
    sectionIndex < statGraphicalSections.length;
    sectionIndex++
  ) {
    while (statGraphicalSections[sectionIndex].firstChild) {
      statGraphicalSections[sectionIndex].removeChild(
        statGraphicalSections[sectionIndex].firstChild
      );
    }
  }
}

function parseStatValue(stat) {
  stat = Math.floor(stat);
  if (stat != 0) return Math.floor(stat);
  return 1;
}

function setColorStatNumbers(num) {
  if (num < 4) return "lightgrey";
  if (num < 7) return "white";
  if (num < 10) return "white";
  if (num < 13) return "white";
  if (num < 16) return "white";
  if (num < 19) return "yellow";
  return "yellow";
}

function createGraphicalStatBars(stats) {
  const statBoxColours = {
    1: "#808080", // Grey
    2: "#7f8186",
    3: "#7e828c",
    4: "#7d8392",
    5: "#7c8498",
    6: "#7b859e",
    7: "#7a86a4",
    8: "#7987aa",
    9: "#7888b0",
    10: "#7789b6",
    11: "#769abc",
    12: "#75abc2",
    13: "#74bcc8",
    14: "#73cdd2",
    15: "#72dee2",
    16: "#71dff2",
    17: "#70e0f2",
    18: "#6fe1f2",
    19: "#6ee2f2",
    20: "#6de3f2", // Very Muted Blue
  };

  const statValues = document.querySelectorAll(".stat-value");
  for (let x = 0; x < statValues.length; x++) {
    statValues[x].innerText = parseStatValue(stats[x]);
    statValues[x].style.color = setColorStatNumbers(statValues[x].innerText);
  }

  const statGraphicalSections = document.querySelectorAll(".stat-graphical");

  for (
    let sectionIndex = 0;
    sectionIndex < statGraphicalSections.length;
    sectionIndex++
  ) {
    for (let x = 0; x < 20; x++) {
      const statLevelBox = document.createElement("div");
      statLevelBox.style.borderTop = "1px solid white";
      statLevelBox.style.borderBottom = "2px solid black";
      statLevelBox.style.flexGrow = "1";
      statLevelBox.style.padding = "4px 6px";
      statLevelBox.style.borderRadius = "0";

      if (x + 1 <= stats[sectionIndex]) {
        if (stats[sectionIndex]) {
        }
        statLevelBox.style.backgroundColor = statBoxColours[x + 1];
      }
      statGraphicalSections[sectionIndex].appendChild(statLevelBox);
    }
  }
}

function createPedigree() {
  const horsePedigreeDiv = document.getElementById("horse-pedigree");
  for (let x = 0; x < 4; x++) {
    const generationDiv = document.createElement("div");
    generationDiv.style.display = "flex";
    generationDiv.style.flexDirection = "column";
    generationDiv.style.flexGrow = "1";
    generationDiv.style.justifyContent = "space-around";
    horsePedigreeDiv.appendChild(generationDiv);
    const amountOfDivs = Math.pow(2, x + 1);
    for (let y = 0; y < amountOfDivs; y++) {
      const newParentDiv = document.createElement("div");
      newParentDiv.style.display = "flex";
      newParentDiv.style.alignItems = "center";
      newParentDiv.style.flexGrow = "1";
      newParentDiv.style.border = "1px solid black"; // For visual debugging
      newParentDiv.style.borderTopLeftRadius = "10px";
      newParentDiv.innerText = "Poo";
      newParentDiv.style.paddingLeft = "0.8rem";
      newParentDiv.style.backgroundColor = "rgb(170, 189, 218)";

      if (y % 2 != 0) {
        newParentDiv.style.backgroundColor = "pink";
      }

      generationDiv.appendChild(newParentDiv);
    }
  }
}
