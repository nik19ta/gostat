import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'

import { cookies } from "next/headers";
import { changeLanguage, defaultLang } from './shared/libs/i18n';
import { APP_LANGUAGES_TYPE } from './shared/constants/languages';

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const cookieStore = cookies();
  const theme = cookieStore.get("theme");
  const lang = (cookieStore.get("lang")?.value ?? defaultLang).toLowerCase() as APP_LANGUAGES_TYPE;

  changeLanguage(lang);

  return (
    <html lang={lang ?? "en"}>
      <body className={`${inter.className} ${theme?.value}`}>
        {children}
      </body>
    </html>
  );
}
