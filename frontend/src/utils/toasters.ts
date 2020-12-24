import { toast, ToastOptions } from 'react-toastify'

const options: ToastOptions = {
  position: toast.POSITION.BOTTOM_RIGHT,
  closeOnClick: true,
  pauseOnHover: false,
  autoClose: 4000
}

export const toastNotification = {
  success: (msg: string) => {
    return toast.success(msg, options)
  },
  error: (errorMsg: string) => {
    return toast.error(errorMsg, options)
  }
}