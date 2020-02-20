let artists;

const cards = document.querySelector('.cards');
const search = document.getElementById('search');


const catchData = (filter) => {

    const url = 'http://localhost:1998/answer'
    fetch(url)
        .then((res) => res.json())
        .then(filter)
        .then(jsx => {
            artists = jsx;
            renderData(artists)
        })

}

catchData();

function renderData(data) {
    cards.innerHTML = ``;
    let counter = 0
    data.forEach((element, ) => {
        let rel = element.Relations.index[0].datesLocations;

        cards.innerHTML += `
        <div style="margin-bottom: 50px;" class="col-md-5" id=${element.id}>
            <div class="card">
            <div class="card__image-container text-center">
                <img class="card__image" width="70%" src="${element.image}" alt="">
            </div>
            
            <svg class="card__svg" viewBox="0 0 800 500">
    
            <path d="M 0 100 Q 50 200 100 250 Q 250 400 350 300 C 400 250 550 150 650 300 Q 750 450 800 400 L 800 500 L 0 500" stroke="transparent" fill="#333"/>
            <path class="card__line" d="M 0 100 Q 50 200 100 250 Q 250 400 350 300 C 400 250 550 150 650 300 Q 750 450 800 400" stroke="pink" stroke-width="5" fill="transparent"/>
            </svg>
        
            <div class="card__content">
                <h2 class="card__title">${element.name}</h2>
                <p>${element.members}</p>
                <p><label><strong>Creation date:</strong></label> ${element.creationDate}</p>
                <p><label><strong>First album:</strong></label> ${element.firstAlbum}</p>
                <table class="table">
                    <thead class="thead-dark">
                        <tr>
                            <th scope="col">    
                                <strong>Locations</strong>
                            </th>
                            <th scope="col">
                                <strong>Dates</strong>
                            </th>
                        </tr>
                    </thead>
                    <tbody id=${'table' + counter}>
                    </tbody>
            </table>
       </div>
       </div>
        `
        let table = document.getElementById('table' + counter);
        let forDoc = 0

        for (let key in rel) {
            table.innerHTML += `
                <tr>
                    <td>${key}</td>
                    <td class="card-dates">${rel[key].map((v) => {
                return v + '<br>';
            }).join("")
                }
                    </td>
                </tr>`

        }
        counter++;
    });

}

const FilterCards = (e) => {
    e.preventDefault();
    let input = e.target;
    catchData(groupNames => groupNames.filter(item => item.name.toLowerCase().includes(input.value.toLowerCase())))
}

search.addEventListener('input', FilterCards);
