import { RxEyeOpen, RxEyeNone } from "react-icons/rx";
import { useState } from "react"

const Input = ({ name, title, placeholder, icon, isPass, shadowLess }) => {

  const [showPass, setShowPass] = useState(false)

  return (
    <label className="flex flex-col gap-3 relative">
      <p className={`font-medium ${!shadowLess ? "text-gray-700" : "text-gray-100"}`}>{title}:</p>
      <div 
        className={`border border-gray-400 ${!shadowLess && "shadow-lg shadow-gray-200"} flex items-center py-2 px-5 gap-5 rounded-md bg-white`}
      >
        <p 
          className="text-lg text-gray-600"
        >
          {icon}
        </p>
        <input 
          type={!showPass && isPass ? "password" : "text"} 
          className={`w-full outline-none ${!shadowLess ? "placeholder:text-gray-300" : "placeholder:text-gray-700"}`} 
          placeholder={placeholder} 
          name={name} 
        />

        {isPass && 
          <div 
            className="absolute right-5" 
            onClick={() => setShowPass(!showPass)}
          >
            {!showPass ? <RxEyeOpen /> : <RxEyeNone />}
          </div>
        }

      </div>
    </label>
  )
}

export default Input