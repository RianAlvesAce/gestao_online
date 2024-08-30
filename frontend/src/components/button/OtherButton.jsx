import { useEffect, useState } from "react"

const OtherButton = ({ title, active, setFunc, setActList }) => {

  const [isActive, setIsActive] = useState(active)

  useEffect(() => {
    setFunc(setIsActive)
  }, [])

  const handleClick = () => {
    setActList.forEach(set => {
      set(false)
    })
    setIsActive(true)
  }

  return (
    <div 
      className={`w-3/4 ${isActive ? "bg-slate-700 text-white border-transparent" : "bg-white text-gray-700 border-gray-700 hover:scale-105"} font-semibold text-center py-4 rounded-md border cursor-pointer transition`}
      onClick={handleClick}
    >
      {title}
    </div>
  )
}

export default OtherButton