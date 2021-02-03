import React, {useState} from 'react';



function App() {
    const [buttonsCount, setCount] = useState(1)

    const [primary, setPrimary] = useState(window.GetColor("primary"))
    const [secondary, setSecondary] = useState(window.GetColor("secondary"))
    const [text, setTextColor] = useState(window.GetColor("text"))

    setInterval(() => {
        setPrimary(window.GetColor("primary"))
        setSecondary(window.GetColor("secondary"))
        setTextColor(window.GetColor("text"))
    }, 333)

    return (
        <div className={"app-barrier"}>
            {
                Array(buttonsCount).fill(1).map(
                    (value, index) =>
                        <button key={index} style={{backgroundColor: primary, color: text, padding: 10 + "px"}}>
                            Some fucking text
                        </button>
                )
            }
            <button onClick={e => setCount(buttonsCount + 1)}>
                Add shiza button
            </button>
        </div>
    );
}

export default App;
