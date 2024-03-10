import './App.css'
import {
  QueryClient,
  QueryClientProvider
} from '@tanstack/react-query'
import { Pages } from '@/components/_pages/pages'
import { store } from './lib/store'
import { Provider } from 'react-redux';
import { enableMapSet } from 'immer';


const queryClient = new QueryClient()


function App() {
  enableMapSet()
  return (
    <Provider store={store}>
      <QueryClientProvider client={queryClient}>
        <Pages/>
      </QueryClientProvider>
    </Provider>
  )
}

export default App