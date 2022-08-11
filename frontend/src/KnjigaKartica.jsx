import React from "react";

const KnjigaKartica = ({knjiga}) =>{
    return (
        <div key = {knjiga.Id}>
            <h2>{knjiga.Naziv}</h2>
        </div>
    )
}

export default KnjigaKartica