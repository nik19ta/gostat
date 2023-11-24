import { SVGProps } from 'react';

export const LightCat = ({ width = 24, height = 24, ...props }: SVGProps<SVGSVGElement>) => (
  <svg xmlns="http://www.w3.org/2000/svg" width="160" height="120" viewBox="0 0 160 120" fill="none" {...props} >
    <path fillRule="evenodd" clipRule="evenodd" d="M90 0H80V10H70V20H60H50H40V10H30V0H20H10V10V20V30H0V40V50V60V70H10V80H20V90V100V110V120H30H40H50H60H70H80H90H100H110H120H130H140V110H150V100H160V90V80V70H150V60H140H130V70V80V90H120H110V80V70V60V50V40V30H100V20V10V0H90Z" fill="#999999"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M30 90H40H50H60H70H80V100H70V110H60V120H50V110H40V120H30V110V100V90Z" fill="white"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M150 70H140V80V90H130V100H120H110V90H100V80H90H80V90H90V100H80V110H90H100H110H120H130H140V100H150V90V80V70Z" fill="white"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M30 10H20V20V30H10V40V50V60V70H20V80H30H40H50H60H70H80V70H90H100V60V50V40V30H90V20V10H80V20H70V30H60H50H40V20H30V10ZM20 60V50H30V60H20ZM60 60H70V50H60V60Z" fill="white"/>
  </svg>
);

export const DarkCat = ({ width = 24, height = 24, ...props }: SVGProps<SVGSVGElement>) => (
  <svg xmlns="http://www.w3.org/2000/svg" width="110" height="83" viewBox="0 0 110 83" fill="none" {...props} >
    <path fillRule="evenodd" clipRule="evenodd" d="M61.875 0H55V6.875H48.125V13.75H41.25H34.375H27.5V6.875H20.625V0H13.75H6.875V6.875V13.75V20.625H0V27.5V34.375V41.25V48.125H6.875V55H13.75V61.875V68.75V75.625V82.5H20.625H27.5H34.375H41.25H48.125H55H61.875H68.75H75.625H82.5H89.375H96.25V75.625H103.125V68.75H110V61.875V55V48.125H103.125V41.25H96.25H89.375V48.125V55V61.875H82.5H75.625V55V48.125V41.25V34.375V27.5V20.625H68.75V13.75V6.875V0H61.875Z" fill="#181818"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M20.625 61.875H27.5H34.375H41.25H48.125H55V68.75H48.125V75.625H41.25V82.5H34.375V75.625H27.5V82.5H20.625V75.625V68.75V61.875Z" fill="#3C3C3C"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M103.125 48.125H96.25V55V61.875H89.375V68.75H82.5H75.625V61.875H68.75V55H61.875H55V61.875H61.875V68.75H55V75.625H61.875H68.75H75.625H82.5H89.375H96.25V68.75H103.125V61.875V55V48.125Z" fill="#3C3C3C"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M20.625 6.875H13.75V13.75V20.625H6.875V27.5V34.375V41.25V48.125H13.75V55H20.625H27.5H34.375H41.25H48.125H55V48.125H61.875H68.75V41.25V34.375V27.5V20.625H61.875V13.75V6.875H55V13.75H48.125V20.625H41.25H34.375H27.5V13.75H20.625V6.875ZM13.75 41.25V34.375H20.625V41.25H13.75ZM41.25 41.25H48.125V34.375H41.25V41.25Z" fill="#3C3C3C"/>
  </svg>
);

export const LightDarkCat = ({ width = 24, height = 24, ...props }: SVGProps<SVGSVGElement>) => (
  <svg xmlns="http://www.w3.org/2000/svg" width="110" height="83" viewBox="0 0 110 83" fill="none" {...props} >
    <path fillRule="evenodd" clipRule="evenodd" d="M61.875 0H55V6.875H48.125V13.75H41.25H34.375H27.5V6.875H20.625V0H13.75H6.875V6.875V13.75V20.625H0V27.5V34.375V41.25V48.125H6.875V55H13.75V61.875V68.75V75.625V82.5H20.625H27.5H34.375H41.25H48.125H55H61.875H68.75H75.625H82.5H89.375H96.25V75.625H103.125V68.75H110V61.875V55V48.125H103.125V41.25H96.25H89.375V48.125V55V61.875H82.5H75.625V55V48.125V41.25V34.375V27.5V20.625H68.75V13.75V6.875V0H61.875Z" fill="#181818"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M13.75 0H20.625V6.875H27.5V13.75H34.375H41.25H48.125V6.875H55H61.875V13.75V20.625H68.75V27.5V34.375V41.25V48.125H61.875H55V55H61.875H68.75V61.875H75.625V68.75H82.5H89.375V61.875H96.25V55V48.125H103.125V55V61.875V68.75H96.25V75.625H89.375H82.5H75.625H68.75H61.875H55V82.5H48.125H41.25H34.375H27.5H20.625H13.75V75.625V68.75V61.875V55H6.875V48.125H0V41.25V34.375V27.5V20.625H6.875V13.75V6.875V0H13.75ZM61.875 68.75H55V61.875H61.875V68.75Z" fill="#999999"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M20.625 61.875H27.5H34.375H41.25H48.125H55V68.75H48.125V75.625H41.25V82.5H34.375V75.625H27.5V82.5H20.625V75.625V68.75V61.875Z" fill="white"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M103.125 48.125H96.25V55V61.875H89.375V68.75H82.5H75.625V61.875H68.75V55H61.875H55V61.875H61.875V68.75H55V75.625H61.875H68.75H75.625H82.5H89.375H96.25V68.75H103.125V61.875V55V48.125Z" fill="#3C3C3C"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M20.625 6.875H13.75V13.75V20.625H6.875V27.5V34.375V41.25V48.125H13.75V55H20.625H27.5H34.375H41.25H48.125H55V48.125H61.875H68.75V41.25V34.375V27.5V20.625H61.875V13.75V6.875H55V13.75H48.125V20.625H41.25H34.375H27.5V13.75H20.625V6.875ZM13.75 41.25V34.375H20.625V41.25H13.75ZM41.25 41.25H48.125V34.375H41.25V41.25Z" fill="white"/>
    <path fillRule="evenodd" clipRule="evenodd" d="M61.875 6.875H55V13.75V20.625V27.5V34.375V41.25V48.125H61.875H68.75V41.25V34.375V27.5V20.625H61.875V13.75V6.875Z" fill="#3C3C3C"/>
  </svg>
);