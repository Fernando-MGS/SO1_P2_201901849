import { Pie } from "react-chartjs-2";

import '../App.css'
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
ChartJS.register(ArcElement, Tooltip, Legend);


const DataUsage = (props) => {
    const dataCpu = {
        labels: ['% Ocupado', '% Libre'],
        datasets: [
            {
                label: '# of Votes',
                data: [props.cpu.occupied, props.cpu.free],
                backgroundColor: [
                    'rgba(230, 6, 6, 1)',
                    '#8fc1ec'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)'
                ],
                borderWidth: 1,
            },
        ],
    };

    const dataRam = {
        labels: ['% Ocupado', '% Libre'],
        datasets: [
            {
                label: '# of Votes',
                data: [props.ram.occupied, props.ram.free],
                backgroundColor: [
                    'rgba(230, 6, 6, 1)',
                    '#8fc1ec'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)'
                ],
                borderWidth: 1,
            },
        ],
    };

    return (
        <div className='dataUsage'>
            <div className="cpu">
                <figure className="figure">
                    <Pie data={dataCpu} />
                    <br/>
                    <figcaption className="figure-caption"><h3>CPU</h3></figcaption>
                </figure>

            </div>
            <div className="ram">
                <figure className="figure">
                    <Pie data={dataRam} />
                    <br/>
                    <figcaption className="figure-caption"><h3>RAM</h3></figcaption>
                </figure>
            </div>
            <hr/>
        </div>
    );
}

export default DataUsage