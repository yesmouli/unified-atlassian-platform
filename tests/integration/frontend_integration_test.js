import { render, screen } from '@testing-library/react';
import App from '../src/App';

test('renders the application', () => {
  render(<App />);
  const linkElement = screen.getByText(/Home/i);
  expect(linkElement).toBeInTheDocument();
});
