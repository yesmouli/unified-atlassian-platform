import { render, screen } from '@testing-library/react';
import JiraTaskBoard from '../src/components/JiraTaskBoard';

test('renders Jira Task Board component', () => {
  render(<JiraTaskBoard />);
  const linkElement = screen.getByText(/Jira Task Board/i);
  expect(linkElement).toBeInTheDocument();
});
