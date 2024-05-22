import '../App.css'

function SearchBarArea(){

    return (
        <div class = 'bg-second'>
            <div class = 'bg-second py-8 px-10 grid justify-items-center grid-cols-3' >
                <div> 
                    <input class = 'py-2 px-1 rounded text-primary focus:outline-none' type='text' placeholder='Fra' list='cityname'/>
                    <datalist id='cityname'>
                        <option value='Oslo S'/>
                        <option value='Trondheim S'/>
                        <option value='Lillehammer stasjon'/>
                        <option value='Stavanger stasjon'/>
                        <option value='Bergen stasjon'/>
                        <option value='Lillestrøm stasjon'/>
                        <option value='Fredrikstad stasjon'/>
                        <option value='Kristiansand stasjon'/>
                        <option value='Kongsvinger stasjon'/>
                    </datalist>   
                </div>
                <div>
                    <button type='button' class=' bg-primary  hover:bg-second hover:ring hover:ring-primary rounded text-white p-2'>
                    Bytt retning               
                    </button>
                </div>

                <div >
                    <input class = 'py-2 px-1 rounded text-primary focus:outline-none' type='text' placeholder='Til' list='cityname'/>
                </div>
            </div>
            <img src='/search_bar_kunst_remix.png'/>
        </div>
    )

}

export default SearchBarArea;