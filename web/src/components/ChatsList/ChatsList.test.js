import { render, screen } from '@testing-library/react';
import ChatsPage from './ChatsList';

xtest('renders learn react link', () => {
  render(<ChatsPage />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
