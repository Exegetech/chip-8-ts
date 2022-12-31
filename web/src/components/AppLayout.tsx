import React from 'react';

function Header(): JSX.Element {
  return (
    <header className="navbar mt-2 mb-2">
      <section className="navbar-section bg-secondary">
        <span className="navbar-brand text-primary">Chip-8 Emulator</span>
      </section>
    </header>
  )
}

function Footer(): JSX.Element {
  return (
    <footer className="bg-secondary text-primary text-center mt-2 mb-2">
      Made by Exegetech
    </footer>
  )
}

export function AppLayout(props): JSX.Element { 
  return (
    <div className="container">
      <Header />

      {props.children}

      <Footer />
    </div>
  );
}

