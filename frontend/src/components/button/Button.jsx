import { useEffect, useState, forwardRef } from "react"

const Button = forwardRef(({ title, icon, active, setFunc, setList, page, setPage }, ref) => {

  const [isActive, setIsActive] = useState(active)

  const handleClick = () => {
    setList.forEach((value) => {
      value(false)
      console.log(title)
    })
    setIsActive(true)
    setPage(page)
  }

  useEffect(() => {
    setFunc(setIsActive)
  }, [])

  return (
    <div 
      className={`${isActive ? "bg-white text-gray-700" : "bg-gray-700 bg-opacity-40 text-white scale-100 hover:scale-105"} flex items-center px-6 py-4 gap-5 rounded-md cursor-pointer transition`} 
      ref={ref}
      onClick={handleClick}
    >
      <p className="text-2xl">
        {icon}
      </p>
      <p className="font-bold">{title}</p>
    </div>
  )
})

export default Button