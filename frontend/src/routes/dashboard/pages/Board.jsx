import Chart from "react-apexcharts"
import PieChart from "../../../components/pieChart/PieChart"
import OtherButton from "../../../components/button/OtherButton"
import { useRef } from "react"
const Board = () => {
  
  const allData = [
    {
      color: "#1460F4",
      label: "N° Atendimentos",
      value: 300
    },
    {
      color: "#56D363",
      label: "Ganhos",
      value: 1000
    }
  ]

  const paymentData = [
    {
      color: "#F4CE14",
      label: "Pix",
      value: 300
    },
    {
      color: "#14F453",
      label: "Cartão",
      value: 300
    },
    {
      color: "#BA56D3",
      label: "Dinheiro",
      value: 300
    },
    {
      color: "#D35656",
      label: "Convênio",
      value: 300
    },
  ]

  const buttonList = [
    {
      title: "Ultima semana",
      active: true
    },
    {
      title: "Ultimo mês"
    },
    {
      title: "Ultimo ano"
    }
  ]

  const btnSetActive = useRef([])

  const addSet = sf => {
    btnSetActive.current.push(sf)
  }

  return (
    <div className="flex-1 flex items-start p-7 gap-10">
      <div className="flex w-max gap-7">

        <PieChart 
          title="Geral"
          data={allData}
        />
        <PieChart 
          title="Formas de pagamento"
          data={paymentData}
        />

      </div>

      <div className="w-1/2">
        <h1 className="mb-7 font-bold text-gray-700">Filtrar por:</h1>
        <div className="flex flex-col gap-4">
          {buttonList.map(button => (
            <OtherButton 
              title={button.title}
              active={button.active}
              setFunc={addSet}
              setActList={btnSetActive.current}
            />
          ))}
        </div>
      </div>
    </div>
  )
}

export default Board