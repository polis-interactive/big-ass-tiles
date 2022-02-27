import React from 'react';
import {TrySendControlValue} from "./api/services";

const controls = [
    "Brightness",
    "Speed",
    "Program",
    "Value"
]

function App() {
    return (
        <div style={{paddingLeft: '1rem'}}>
            <h1>Controls</h1>
            {controls.map((c, i) => (
                <div key={`${c}-control`} >
                    <h2>{c}</h2>
                    <input
                        type='range' min='0' max='1000' defaultValue='0'
                        onChange={(e) => {
                            const actualVal: number = parseInt(e.target.value) / 1000
                            TrySendControlValue(i, actualVal).then()
                        }}
                    />
                </div>
            ))}
    </div>
  );
}

export default App;
