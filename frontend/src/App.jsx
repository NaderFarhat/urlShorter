import { useState } from 'react'
import styled, { ThemeProvider } from 'styled-components'
import { GlobalStyles } from './styles/GlobalStyles'
import { theme } from './styles/colors'
import UrlShortener from './components/UrlShortener'

const AppContainer = styled.div`
  min-height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
  background-color: ${props => props.theme.colors.onyx};
`

const Title = styled.h1`
  color: ${props => props.theme.colors.onyx};
  font-size: 2.5rem;
  margin-bottom: 2rem;
  text-align: center;
`

function App() {
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyles />
      <AppContainer>
        <UrlShortener />
      </AppContainer>
    </ThemeProvider>
  )
}

export default App
