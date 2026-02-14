import { useState } from 'react'
import styled from 'styled-components'

const Container = styled.div`
  width: 100%;
  max-width: 600px;
  background-color: ${props => props.theme.colors.mercury};
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
`

const InputContainer = styled.div`
  display: flex;
  gap: 12px;
  margin-bottom: 2rem;
`

const Input = styled.input`
  flex: 1;
  padding: 12px 16px;
  border: 2px solid ${props => props.theme.colors.twilightGrey};
  border-radius: 8px;
  font-size: 16px;
  background-color: ${props => props.theme.colors.snow};
  color: ${props => props.theme.colors.onyx};
  transition: border-color 0.2s ease;

  &:focus {
    border-color: ${props => props.theme.colors.deepSaffron};
  }

  &::placeholder {
    color: ${props => props.theme.colors.twilightGrey};
  }
`

const Button = styled.button`
  padding: 12px 24px;
  background-color: ${props => props.theme.colors.deepSaffron};
  color: ${props => props.theme.colors.snow};
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  transition: all 0.2s ease;

  &:hover {
    background-color: #e68a1f;
    transform: translateY(-1px);
  }

  &:active {
    transform: translateY(0);
  }

  &:disabled {
    background-color: ${props => props.theme.colors.twilightGrey};
    cursor: not-allowed;
    transform: none;
  }
`

const UrlList = styled.div`
  display: flex;
  flex-direction: column;
  gap: 12px;
`

const UrlItem = styled.div`
  background-color: ${props => props.theme.colors.snow};
  border: 1px solid ${props => props.theme.colors.twilightGrey};
  border-radius: 8px;
  padding: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: border-color 0.2s ease;

  &:hover {
    border-color: ${props => props.theme.colors.deepSaffron};
  }
`

const UrlInfo = styled.div`
  flex: 1;
`

const OriginalUrl = styled.div`
  font-size: 14px;
  color: ${props => props.theme.colors.liquorice};
  margin-bottom: 4px;
  word-break: break-all;
`

const ShortUrl = styled.div`
  font-size: 16px;
  color: ${props => props.theme.colors.deepSaffron};
  font-weight: 600;
`

const CopyButton = styled.button`
  padding: 8px 16px;
  background-color: ${props => props.theme.colors.liquorice};
  color: ${props => props.theme.colors.snow};
  border-radius: 6px;
  font-size: 14px;
  transition: all 0.2s ease;

  &:hover {
    background-color: ${props => props.theme.colors.onyx};
  }

  &:active {
    transform: scale(0.95);
  }
`

const ErrorMessage = styled.div`
  color: #e74c3c;
  font-size: 14px;
  margin-top: 8px;
`

const EmptyState = styled.div`
  text-align: center;
  color: ${props => props.theme.colors.twilightGrey};
  padding: 2rem;
`

const UrlShortener = () => {
  const [url, setUrl] = useState('')
  const [urls, setUrls] = useState([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const isValidUrl = (string) => {
    try {
      new URL(string)
      return true
    } catch (_) {
      return false
    }
  }

  const generateShortUrl = () => {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    let result = ''
    for (let i = 0; i < 6; i++) {
      result += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    return result
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')

    if (!url.trim()) {
      setError('Please enter a URL')
      return
    }

    if (!isValidUrl(url)) {
      setError('Please enter a valid URL')
      return
    }

    setIsLoading(true)

    // Simulate API call
    setTimeout(() => {
      const shortCode = generateShortUrl()
      const newUrl = {
        id: Date.now(),
        originalUrl: url,
        shortUrl: `https://short.ly/${shortCode}`,
        createdAt: new Date().toISOString()
      }
      
      setUrls(prev => [newUrl, ...prev])
      setUrl('')
      setIsLoading(false)
    }, 500)
  }

  const copyToClipboard = async (text) => {
    try {
      await navigator.clipboard.writeText(text)
      // You could add a toast notification here
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }

  return (
    <Container>
      <form onSubmit={handleSubmit}>
        <InputContainer>
          <Input
            type="text"
            placeholder="Enter your long URL..."
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            disabled={isLoading}
          />
          <Button type="submit" disabled={isLoading}>
            {isLoading ? 'Shortening...' : 'Shorten'}
          </Button>
        </InputContainer>
        {error && <ErrorMessage>{error}</ErrorMessage>}
      </form>

      <UrlList>
        {urls.length === 0 ? (
          <EmptyState>
            No shortened URLs yet. Create your first one above!
          </EmptyState>
        ) : (
          urls.map((urlItem) => (
            <UrlItem key={urlItem.id}>
              <UrlInfo>
                <OriginalUrl>{urlItem.originalUrl}</OriginalUrl>
                <ShortUrl>{urlItem.shortUrl}</ShortUrl>
              </UrlInfo>
              <CopyButton onClick={() => copyToClipboard(urlItem.shortUrl)}>
                Copy
              </CopyButton>
            </UrlItem>
          ))
        )}
      </UrlList>
    </Container>
  )
}

export default UrlShortener
