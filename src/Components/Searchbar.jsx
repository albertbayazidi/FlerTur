import '../App.css'

function SearchBarArea(){

    return (
        <div class = 'bg-second'>
            <div class = 'bg-second p-10 grid justify-items-center grid-cols-3' >
                <div> 
                    <input class = 'py-2 rounded px-1 text-black focus:outline-none' type='text' placeholder='Fra'/>     
                </div>
                <div>
                    <button type='button' class=' bg-primary  hover:bg-second hover:ring hover:ring-primary rounded text-white p-2'>
                    Bytt retning               
                    </button>
                </div>

                <div >
                    <input class = 'py-2 rounded px-1 text-black focus:outline-none' type='text' placeholder='Til'/>
                </div>
            </div>
            <div class = 'grid grid-cols-3 justify-items-center'> 
                <div>
                    <input class = ' mb-1 py-2 px-1 rounded w-24 focus:outline-none' type = 'text' placeholder='Dato'/>
                    
                    <button type='button' class='mb-1 ml-5 p-2  bg-primary hover:bg-second hover:ring hover:ring-primary rounded text-white '>
                    Søk
                    </button>
                </div>
            </div>
            <img src='/search_bar_kunst_remix.png'/>
        </div>
    )

}

export default SearchBarArea;