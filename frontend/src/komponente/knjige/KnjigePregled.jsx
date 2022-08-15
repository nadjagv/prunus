import React, {useEffect, useState} from "react";
import KnjigaKartica from "./KnjigaKartica";
import {useNavigate} from "react-router-dom"

const GATEWAY_URL = "http://localhost:8080/knjige"

const KnjigePregled = () =>{
    const [knjige, setKnjige] = useState([])

    const navigate = useNavigate()

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
                <div className = "container">
                    {knjige.map((knjiga) => (
                        <div onClick={() => navigate("/knjige/" + knjiga.Id)}>
                            <KnjigaKartica knjiga={knjiga} key = {knjiga.Isbn}/>
                        </div>
                    ))}
                </div>
            ):(
                <div className = "empty">
                    <h2>Nema knjiga za pregled.</h2>
                </div>
            )}

        </div>
    )
}

export default KnjigePregled