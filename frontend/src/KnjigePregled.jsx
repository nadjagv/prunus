import React, {useEffect, useState} from "react";
import KnjigaKartica from "./KnjigaKartica";

const GATEWAY_URL = "http://localhost:8080/knjige"

const KnjigePregled = () =>{
    const [knjige, setKnjige] = useState([])

    useEffect(()=>{
        preuzmiSve()
        
    }, [])

    const preuzmiSve = async () => {
        const response = await fetch (`${GATEWAY_URL}`)
        const data = await response.json();
        console.log(data)
        setKnjige(data)
    }

    return (
        <div>

            {knjige.length > 0 ? (
                <div>
                    {knjige.map((knjiga) => (
                        <KnjigaKartica knjiga={knjiga}/>
                    ))}
                </div>
            ):(
                <div>
                    <h2>Nema knjiga za pregled.</h2>
                </div>
            )}

        </div>
    )
}

export default KnjigePregled