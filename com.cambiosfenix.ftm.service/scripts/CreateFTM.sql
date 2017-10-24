﻿USE [FTM]
GO
/****** Object:  Database [prm]    Script Date: 20/10/2017 03:05:37 p.m. ******/
CREATE DATABASE [FTM]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'ftm', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL12.SQLEXPRESS\MSSQL\DATA\dtm.mdf' , SIZE = 5120KB , MAXSIZE = UNLIMITED, FILEGROWTH = 1024KB )
 LOG ON 
( NAME = N'ftm_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL12.SQLEXPRESS\MSSQL\DATA\ftm_log.ldf' , SIZE = 2048KB , MAXSIZE = 2048GB , FILEGROWTH = 10%)
GO
ALTER DATABASE [FTM] SET COMPATIBILITY_LEVEL = 120
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [FTM].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [FTM] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [FTM] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [FTM] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [FTM] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [FTM] SET ARITHABORT OFF 
GO
ALTER DATABASE [FTM] SET AUTO_CLOSE OFF 
GO
ALTER DATABASE [FTM] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [FTM] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [FTM] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [FTM] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [FTM] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [FTM] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [FTM] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [FTM] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [FTM] SET  DISABLE_BROKER 
GO
ALTER DATABASE [FTM] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [FTM] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [FTM] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [FTM] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [FTM] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [FTM] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [FTM] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [FTM] SET RECOVERY SIMPLE 
GO
ALTER DATABASE [FTM] SET  MULTI_USER 
GO
ALTER DATABASE [FTM] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [FTM] SET DB_CHAINING OFF 
GO
ALTER DATABASE [FTM] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [FTM] SET TARGET_RECOVERY_TIME = 0 SECONDS 
GO
ALTER DATABASE [FTM] SET DELAYED_DURABILITY = DISABLED 
GO
USE [FTM]
GO
/****** Object:  Table [dbo].[Transaction]    Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Transaction](
	[TransactionID] [int] IDENTITY(1,1) NOT NULL,
	[TransactionCode] [nvarchar](70) NOT NULL,
	[Type] [int] NOT NULL,
	[Date] [date] NULL,
	[TRMID] [int] NOT NULL,
	[TRMValue] [decimal] (18,7) NULL,
	[PartnerID] [int] NOT NULL,
	[CurrencyFromQuantity] [float]  NULL,
	[CurrencyToQuantity] [float]  NULL,
	[PayTypeID] [int] NOT NULL,
	[IncoTermID] [int] NOT NULL,
	[Casher] [nvarchar](50)  NULL,
	[Notes] [nvarchar](70)  NULL,
 CONSTRAINT [PK_Transaction] PRIMARY KEY CLUSTERED 
(
	[TransactionID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Partner]    Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Partner](
	[PartnerID] [int] IDENTITY(1,1) NOT NULL,
	[PartnerCode] [nvarchar](70) NOT NULL,
	[Name] [nvarchar](140)  NULL,
	[Address] [nvarchar](140)  NULL,
	[City] [nvarchar](30)  NULL,
	[ZipCode] [nvarchar](70)  NULL,
	[State] [nvarchar](30)  NULL,
	[CountryCode] [nvarchar](10)  NULL,
	[Phone] [nvarchar](30)  NULL,
	[Email] [nvarchar](30)  NULL,
 CONSTRAINT [PK_Partner] PRIMARY KEY CLUSTERED 
(
	[PartnerID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Currency]    Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Currency](
	[CurrencyID] [int] IDENTITY(1,1) NOT NULL,
	[CurrencyCode] [nvarchar](70) NOT NULL,
	[Description] [nvarchar](70)  NULL,
 CONSTRAINT [PK_Currency] PRIMARY KEY CLUSTERED 
(
	[CurrencyID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[TRM]    Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[TRM](
	[TRMID]  [int] IDENTITY(1,1) NOT NULL,
	[CurrencyFromID] [int] NOT NULL,
	[CurrencyToID] [int] NOT NULL,
	[TRMValue] [decimal] (18,7) NULL,
	[TRMSaleValue] [decimal] (18,7) NULL,
 CONSTRAINT [PK_TRM] PRIMARY KEY CLUSTERED 
(
	[TRMID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[EnumerationSet]    Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[EnumerationSet](
	[EnumerationSetID] [int] IDENTITY(1,1) NOT NULL,
	[EnumerationSetCode] [nvarchar](70)  NULL,
	[Description] [nvarchar](70)  NULL,
 CONSTRAINT [PK_EnumerationSet] PRIMARY KEY CLUSTERED 
(
	[EnumerationSetID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Enumeration]   Script Date: 20/10/2017 03:05:37 p.m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Enumeration](
	[EnumerationID] [int] IDENTITY(1,1) NOT NULL,
	[EnumerationSetID] [int] NOT NULL,
	[EnumerationValue] [int]   NULL,
	[EnumerationText]  [nvarchar](70)  NULL,
 CONSTRAINT [PK_Enumeration] PRIMARY KEY CLUSTERED 
(
	[EnumerationID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO

--SET IDENTITY_INSERT [dbo].[Partner] ON 
INSERT [dbo].[Partner] ([PartnerCode], [Name], [Address], [City], [ZipCode], [State], [CountryCode], [Phone], [Email]) VALUES (N'000001',N'Promotora de Turismo',N'Calle 9 # 48-03',N'Cali',N'763600',N'Valle del Cauca',N'CO',N'3166874526',N'turismo@email.com')
INSERT [dbo].[Partner] ([PartnerCode], [Name], [Address], [City], [ZipCode], [State], [CountryCode], [Phone], [Email]) VALUES (N'000002',N'Javeriana',N'Calle 100 # 48-03',N'Cali',N'763607',N'Valle del Cauca',N'CO',N'3186878926',N'javeriana@email.com')
--SET IDENTITY_INSERT [dbo].[Partner] OFF

--SET IDENTITY_INSERT [dbo].[Currency] ON 
INSERT [dbo].[Currency] ([CurrencyCode], [Description] ) VALUES ( N'COP', N'Peso Colombiano')
INSERT [dbo].[Currency] ([CurrencyCode], [Description] ) VALUES ( N'USD', N'Dolar Americano')
INSERT [dbo].[Currency] ([CurrencyCode], [Description] ) VALUES ( N'EUR', N'Euro')
--SET IDENTITY_INSERT [dbo].[Currency] OFF

--SET IDENTITY_INSERT [dbo].[TRM] ON 
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 1,1,CAST(1 AS decimal(18, 7)),CAST(1 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 2,2,CAST(1 AS decimal(18, 7)),CAST(1 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 2,1,CAST(2932.55132 AS decimal(18, 7)),CAST(2990.55132 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 2,3,CAST(0.8485038 AS decimal(18, 7)),CAST(0.9 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 3,3,CAST(1 AS decimal(18, 7)),CAST(1 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 3,1,CAST(3456.1437 AS decimal(18, 7)),CAST(3506.1437 AS decimal(18, 7)))
INSERT [dbo].[TRM] ([CurrencyFromID], [CurrencyToID], [TRMValue], [TRMSaleValue]) VALUES ( 3,2,CAST(1.178545  AS decimal(18, 7)),CAST(1.21 AS decimal(18, 7)))
--SET IDENTITY_INSERT [dbo].[TRM] OFF

--SET IDENTITY_INSERT [dbo].[EnumerationSet] ON 
INSERT [dbo].[EnumerationSet] ([EnumerationSetCode],[Description]) VALUES (N'TransactionTypes', N'Types of transactions')
INSERT [dbo].[EnumerationSet] ([EnumerationSetCode],[Description]) VALUES (N'Incoterms', N'List of incoterms')
INSERT [dbo].[EnumerationSet] ([EnumerationSetCode],[Description]) VALUES (N'PayType', N'Forms of pay')
--SET IDENTITY_INSERT [dbo].[EnumerationSet] OFF

--SET IDENTITY_INSERT [dbo].[Enumeration] ON 
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (1,1, N'Compra')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (1,2, N'Venta')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (2,1, N'En Local (EXW)')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (2,2, N'A Domicilio (DAP)')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (3,1, N'Efectivo')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (3,2, N'Cheque')
INSERT [dbo].[Enumeration] ([EnumerationSetID],[EnumerationValue],[EnumerationText]) VALUES (3,3, N'Deposito')
--SET IDENTITY_INSERT [dbo].[Enumeration] OFF

--SET IDENTITY_INSERT [dbo].[Transaction] ON 
INSERT [dbo].[Transaction] ([TransactionCode], [Type], [Date], [TRMID], [TRMValue], [PartnerID],[CurrencyFromQuantity],[CurrencyToQuantity],[PayTypeID],[IncoTermID],[Casher],  [Notes] ) VALUES ( N'002701',1, CAST(N'2017-10-21' AS Date),3, CAST(2945 AS decimal(18, 7)), 1,500, 1472500,1,1, N'JP', N'General Transaction ok')
INSERT [dbo].[Transaction] ([TransactionCode], [Type], [Date], [TRMID], [TRMValue], [PartnerID],[CurrencyFromQuantity],[CurrencyToQuantity],[PayTypeID],[IncoTermID],[Casher],  [Notes] ) VALUES ( N'002702',2, CAST(N'2017-10-22' AS Date),3, CAST(2945 AS decimal(18, 7)), 2,250, 736250,1,1, N'JP', N'General Transaction ok')
--SET IDENTITY_INSERT [dbo].[Transaction] OFF

--FK TRANSACTIONS
ALTER TABLE [dbo].[Transaction]  WITH CHECK ADD  CONSTRAINT [FK_TransactionTRM] FOREIGN KEY([TRMID])
REFERENCES [dbo].[TRM] ([TRMID])
GO
ALTER TABLE [dbo].[Transaction] CHECK CONSTRAINT [FK_TransactionCurrencyFrom]
GO

ALTER TABLE [dbo].[Transaction]  WITH CHECK ADD  CONSTRAINT [FK_TransactionPartner] FOREIGN KEY([PartnerID])
REFERENCES [dbo].[Partner] ([PartnerID])
GO
ALTER TABLE [dbo].[Transaction] CHECK CONSTRAINT [FK_TransactionPartner]
GO
--FK TRM

ALTER TABLE [dbo].[TRM]  WITH CHECK ADD  CONSTRAINT [FK_TRMCurrencyFrom] FOREIGN KEY([CurrencyFromID])
REFERENCES [dbo].[Currency] ([CurrencyID])
GO
ALTER TABLE [dbo].[TRM] CHECK CONSTRAINT [FK_TRMCurrencyFrom]
GO

ALTER TABLE [dbo].[TRM]  WITH CHECK ADD  CONSTRAINT [FK_TRMCurrencyTo] FOREIGN KEY([CurrencyToID])
REFERENCES [dbo].[Currency] ([CurrencyID])
GO
ALTER TABLE [dbo].[TRM] CHECK CONSTRAINT [FK_TRMCurrencyTo]
GO

--FK [Enumeration]
ALTER TABLE [dbo].[Enumeration]  WITH CHECK ADD  CONSTRAINT [FK_EnumerationEnumerationSet] FOREIGN KEY([EnumerationSetID])
REFERENCES [dbo].[EnumerationSet] ([EnumerationSetID])
GO
ALTER TABLE [dbo].[Enumeration] CHECK CONSTRAINT [FK_EnumerationEnumerationSet]
GO
USE [FTM]
GO
ALTER DATABASE [FTM] SET  READ_WRITE 
GO
