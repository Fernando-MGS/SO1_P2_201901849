import axios from 'axios'

export const getData = async () => {
    return await axios.get('http://34.125.39.227:4000/readModCPU')

}