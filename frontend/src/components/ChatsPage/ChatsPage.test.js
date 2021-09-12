import { render, screen } from '@testing-library/react';
import ChatsPage from './ChatsPage';

xtest('renders learn react link', () => {
  render(<ChatsPage />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
