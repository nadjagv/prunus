import React from "react";
import Zanr from "../../enumeracije/Zanr";

const KnjigaKartica = ({knjiga}) =>{
    return (
        <div className="knjiga">
            <div>
                <p>{Object.keys(Zanr).find(key => Zanr[key] === knjiga.Zanr)}</p>
            </div>

            <div>
                <img src={knjiga.Slika} alt={knjiga.Naziv} />
            </div>

            <div>
                <span>{knjiga.ImeAutora} {knjiga.PrezimeAutora}</span>
                <h3>{knjiga.Naziv}</h3>
            </div>
        </div>
    )
}

export default KnjigaKartica