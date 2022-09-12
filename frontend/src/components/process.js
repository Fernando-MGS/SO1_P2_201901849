
const Process = (props) => {

    const userName = (id) => {
        if (id === '1001') {
            return 'fernandogs2002'
        } else if (id === '103') {
            return 'message+'
        } else if (id === '104') {
            return 'syslog'
        } else if (id === '112') {
            return '_chrony'
        } else if (id === '101') {
            return 'systemd+'
        } else if (id === '0') {
            return 'root'
        } else {
            return 'daemon'
        }
    }

    return (
        <div className="container">
            <table className="table table-hover table-dark">
                <thead>
                    <tr>
                        <th scope="col">PID</th>
                        <th scope="col">Nombre</th>
                        <th scope="col">Estado</th>
                        <th scope="col">Propietario</th>
                    </tr>
                </thead>
                <tbody>
                    {                        
                        props.process.map(proc => {
                            return (
                                <tr key={proc.pid}>
                                    <td>{proc.pid}</td>
                                    <td>{proc.comm}</td>
                                    <td>{proc.state}</td>
                                    <td>{userName(proc.owner)} || {proc.owner}</td>
                                </tr>)
                        }
                        )
                    }
                </tbody>
            </table>
        </div>
    )
}

export default Process;