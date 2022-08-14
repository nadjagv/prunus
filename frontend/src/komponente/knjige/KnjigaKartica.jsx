import React from "react";

const KnjigaKartica = ({knjiga}) =>{
    return (
        <div key = {knjiga.Id} className="knjiga">
            <div>{knjiga.Naziv}</div>
        </div>
    )
}

export default KnjigaKartica