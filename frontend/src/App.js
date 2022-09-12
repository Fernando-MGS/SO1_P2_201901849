import './App.css';
import DataUsage from './components/usages';
import Process from './components/process';
import { useEffect, useState } from 'react';
import { getData } from './components/appServices';
import "bootstrap/dist/css/bootstrap.min.css";

function App() {

  const [process, setProcess] = useState([])
  const [summ, setSumm] = useState({ exec: 0, susp: 0, stop: 0, zombie: 0 })
  const [ramGraph, setRamGraph] = useState({ free: 100, occupied: 0 })
  const [cpuGraph, setCpuGraph] = useState({ free: 100, occupied: 0 })
  const loadData = async () => {
    const res = await getData()
    setProcess(res.data.process)
    const cpu = res.data.cpu
    const ram = res.data.ram
    setCpuGraph({ free: cpu.free, occupied: cpu.occupied })
    setRamGraph({ free: ram.free, occupied: ram.occupied })
    console.log(process)
    console.log(res.data)
  }

  const execProc = () => {
    let stop = 0
    let exec = 0
    let susp = 0
    let zombie = 0
    for (let proc of process) {
      if (proc.state === '1') {
        console.log(proc.state)
        exec++
      } else {
        exec++
      }
    }
    setSumm({ exec: stop, susp: susp, stop: stop, zombie: zombie })
  }


  useEffect(() => {
    loadData()
  }, [])

  return (
    <div className="App">
      <DataUsage cpu={cpuGraph} ram={ramGraph} />
      <div className='generalData'>
        <h5>En ejecucion: {summ.exec}</h5>
        <h5>Suspendidos: {summ.susp}</h5>
        <h5>Detenidos: {summ.stop}</h5>
        <h5>Zombie: {summ.zombie}</h5>
        <h5>Total: {process.length}</h5>
      </div>
      <Process process={process} />
    </div>
  );
}

export default App;
