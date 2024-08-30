import { useEffect, useState } from "react"
import Chart from "react-apexcharts"

const PieChart = ({ title, data }) => {

  const [colors, setColors] = useState([])
  const [labels, setLabels] = useState([])
  const [values, setValues] = useState([])

  useEffect(() => {
    if(data) {
      data.forEach((value) => {
        setColors((current) => {
          return [...current, value.color]
        })
        setLabels((current) => {
          return [...current, value.label]
        })
        setValues(current => {
          return [...current, value.value]
        })
      })
    }
  }, [])

  const chart = {
    series: values,
      pie: {
        dataLabels: {
        position: "top",
      }
    },
    xaxis: {
      labels:{
        show: false
      }
    },
    yaxis: {
      labels: {
        show: false
      }
    },
    options: {
      colors: colors,
      legend: {
        show: false,
      },
      dataLabels: {
        enabled: false
      },
      chart: {
        type: "donut"
      },
      labels: labels,
      responsive: [
        {
          legend: {
            show: false
          },
          width: "100"
        }
      ],
    },
  }

  return (
    <div className="flex-1 border border-gray-600 rounded-md shadow-md shadow-gray-300 flex flex-col items-center gap-8 pt-5 text-gray-700">

      <h1 className="font-bold text-xl">{title}</h1>

      <div className="flex flex-wrap gap-x-5 gap-y-2 justify-center px-2">
        {colors && colors.map((color, index) => (
            <div className="flex items-center text-xs gap-2 text-gray-600 font-bold">
              <div className="w-6 h-3" style={{backgroundColor: color}}></div>
              <p>{labels[index]}</p>
            </div>
          ))}
      </div>


      <div className="flex-1 flex items-end pb-5">
        <Chart 
          series={chart.series}
          type="donut"
          options={chart.options}
          width={250}
        />
      </div>
    </div>
  )
}

export default PieChart