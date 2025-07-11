# MonitorV2 Architecture

## Overview

We want to have a powerful tool like [atop](https://www.atoptool.nl/)
for monitoring the performance of EVM based blockchains. Here are the
features:

- Live monitoring of the tip of the chain showing recent blocks
- A block view that shows details of a particular block
- A raw block view that shows the pretty printed json view
- A transaction view that shows the details of a particular
  transaction
- A generic "search" feature that takes as input a particular block
  number, block hash, or transaction hash and will jump directly to
  that view (if it's found)
- A high level overview of the chain
  - Chain ID
  - Gas Prices
- A deep info view that prints a snapshot of all of the information
  that we have. This could be very useful for rollups where there are
  specific system contracts or addresses or extra RPCs that can
  provide additional context
- Event and function signature decoding powered by 4byte.directory
- Multiple renderers - The default is a TUI, but the data can also be
  rendered as a JSON stream
- Reorg detection - If a block hash changes, we can rewind and update
  the store. The depth to check for reorgs is configurable
- TUI support for mouse clicking
- Sortable columns for adjusting the view based on the currently
  indexed blocks

## Design Principles

- Non-blocking design. I.e. when something is loading, the UI remains
  responsive
- Configurable - Many aspects of the tool can be configured
- Unified data access - Single interface for all chain-related data
- Intelligent caching - Different TTL strategies for different data types
- Capability awareness - Graceful handling of unsupported RPC methods

## Architecture Components

### ChainStore (Previously Store)

The **ChainStore** is a unified interface that abstracts away the details of how 
blockchain data is accessed and cached. It consolidates both block data and 
chain metadata into a single coherent interface.

Key features:
- **Unified Interface**: Replaces the previous BlockStore with comprehensive chain data access
- **Intelligent Caching**: TTL-based caching with different strategies:
  - Static data (ChainID): Cached indefinitely
  - Semi-static (Safe/Finalized blocks): 5 minutes TTL
  - Frequent (Gas price, Fee history): 30 seconds TTL
  - Very frequent (Pending/Queued txs): 5 seconds TTL
  - Block-aligned (Base fee): Cached per block
- **Capability Detection**: Automatically tests RPC methods and gracefully handles unsupported endpoints
- **Configurable TTL**: Different cache expiration strategies for different data types

Store implementations:
- **PassthroughStore**: Direct RPC passthrough with intelligent caching (current implementation)
- **Future stores**: SQLite, memory-based, hybrid approaches

### Indexer

The **Indexer** is responsible for fetching blockchain data and coordinating between
the ChainStore and Renderers. It provides a clean abstraction layer and manages
the flow of data through the system.

Key responsibilities:
- **Block fetching**: Polls for new blocks and publishes them to renderers
- **Gap detection**: Identifies and handles missing blocks
- **Parallel processing**: Concurrent block fetching for improved performance
- **Delegation**: Provides unified access to all ChainStore methods
- **Channel-based communication**: Non-blocking data flow to renderers

### Renderer

The **Renderer** interface supports multiple output formats:

#### TviewRenderer (TUI)
- **Page-based architecture**: Home, Block Detail, Transaction Detail, Info, Help pages
- **Dual-pane home layout**: Status pane (1/3) and metrics pane (2/3)
- **Comprehensive status pane**: Real-time chain information including:
  - Current timestamp (full date-time for screenshots)
  - RPC endpoint URL (for network identification)
  - Chain ID, gas prices, pending transactions
  - Safe/finalized block numbers, base fees
- **Advanced metrics pane**: 8-row atop-style metrics display:
  - **BLOK**: Latest, safe, and finalized block numbers
  - **THRU**: Transaction and gas throughput (10s/30s averages)
  - **GAS**: Base fee averages and current gas price
  - **POOL**: Pending and queued transaction counts
  - **SIG1**: EOA transactions and contract deployments
  - **SIG2**: ERC20 transfers and NFT transactions
  - **ACCO**: Unique from/to address counters
  - **PEER**: Network peer count
- **Enhanced blocks table**: 10-column display with:
  - Block number, absolute/relative time, block intervals
  - Transaction count, human-readable size
  - Gas usage with percentage and formatted numbers
  - State root hash
- **Interactive navigation**: 
  - Block detail view with side-by-side transaction table and raw JSON
  - Transaction detail view with human-readable properties and receipt JSON
  - Breadcrumb-style navigation with Escape key
  - Tab navigation between panes
- **Modal system**: Focus-protected quit dialog with background update resistance
- **Keyboard shortcuts**: Intuitive navigation with h/i/q/Esc keys
- **Real-time updates**: Multiple update cycles (5-15 second intervals)

#### JSONRenderer
- Structured JSON output for automation and scripting

### Metrics System

The advanced metrics system provides real-time blockchain analytics in an atop-style format:

#### Transaction Analysis
- **Method Signature Detection**: String-based matching for ERC20/NFT method selectors
- **Transaction Classification**: EOA transfers vs contract deployments vs contract calls
- **Performance Optimized**: Direct string prefix matching eliminates hex decoding overhead
- **Real-time Counters**: Continuously updated across all loaded blocks

#### Throughput Calculations
- **Windowed Averages**: 10-second and 30-second rolling calculations
- **TPS (Transactions Per Second)**: Transaction throughput metrics
- **GPS (Gas Per Second)**: Gas consumption rate metrics
- **Base Fee Tracking**: Average base fees over time windows

#### Address Analytics
- **Unique Address Tracking**: Distinct from/to addresses across all transactions
- **Contract Creation Filtering**: Excludes zero-address transactions appropriately
- **Memory Efficient**: Hash-map based uniqueness detection

### Data Types and Formatting

The system includes sophisticated data formatting for human readability:
- **Relative timestamps**: "5m ago", "2h ago", "3d ago"
- **Byte sizes**: "1.5KB", "2.3MB" with proper units
- **Number formatting**: Comma-separated thousands (e.g., "13,402,300")
- **Gas percentages**: Utilization display (e.g., "44.7%")
- **Hash truncation**: Smart abbreviation with ellipsis
- **Throughput formatting**: Scientific notation for large values (1.2M, 3.4K)

## Current Implementation Status

### Completed Features
- Unified ChainStore architecture replacing BlockStore
- Intelligent caching system with TTL strategies
- RPC capability detection and graceful error handling
- Page-based TUI with comprehensive keyboard navigation
- Advanced dual-pane home layout (status + metrics)
- 8-row atop-style metrics system with real-time counters
- 10-column enhanced blocks table with sortable columns
- Interactive block detail page with side-by-side panes
- Transaction detail page with human-readable formatting
- Async transaction signature lookup via 4byte.directory
- Modal focus management system (quit dialog, future search)
- Breadcrumb navigation with Escape key support
- Multiple renderer support (JSON, TUI)
- Human-readable data formatting utilities
- Performance-optimized transaction counters

### Recently Implemented
- **Transaction Analysis**: EOA, contract deployment, ERC20, and NFT counters
- **Address Tracking**: Unique from/to address counters across all transactions
- **Throughput Metrics**: Real-time TPS/GPS calculations with windowed averages
- **UI Navigation**: Breadcrumb-style page switching with focus management
- **Modal System**: Focus-protected dialogs that resist background update interference
- **Signature Decoding**: Integration with 4byte.directory for method/event identification

### Future Features
- Search functionality (block/tx lookup)
- Raw JSON block view
- Mouse support for TUI
- Reorg detection and handling
- Event/function signature decoding via 4byte.directory
- Deep info view for rollup-specific data
- Additional store implementations (SQLite, memory)

## Data Flow

```
┌─────────┐    ┌──────────────┐    ┌─────────┐    ┌───────────┐    ┌──────────┐
│   RPC   │ -> │  ChainStore  │ -> │ Indexer │ -> │  Metrics  │ -> │ Renderer │
│ Network │    │   (Cached)   │    │         │    │  Engine   │    │   (TUI)  │
└─────────┘    └──────────────┘    └─────────┘    └───────────┘    └──────────┘
                                        │              │
                                        v              v
                                   ┌─────────┐    ┌─────────┐
                                   │ Blocks  │    │ Real-   │
                                   │ Channel │    │ time    │
                                   │         │    │ Metrics │
                                   └─────────┘    └─────────┘
```

1. **RPC Network**: Source of truth for all blockchain data
2. **ChainStore**: Unified data access with intelligent caching and capability detection
3. **Indexer**: Coordinates data flow, provides block channels, and manages concurrent fetching
4. **Metrics Engine**: Real-time analytics processing with throughput, counters, and address tracking
5. **Renderer**: Presents data in user-friendly format with advanced metrics display (TUI, JSON, etc.)

### Multi-Channel Architecture

The system uses multiple data channels for responsive updates:
- **Block Channel**: New block notifications for table updates
- **Metrics Channel**: Real-time analytics updates for metrics pane
- **Background Updates**: Periodic chain info, network status, and block info refreshes

### Caching Strategy

The ChainStore implements a sophisticated caching strategy based on data characteristics:

- **Static data** (ChainID): Never expires, fetched once
- **Semi-static data** (Safe/Finalized blocks): 5-minute TTL
- **Frequent data** (Gas prices, Fee history): 30-second TTL  
- **Very frequent data** (Pending transactions): 5-second TTL
- **Block-aligned data** (Base fees): Cached per block number

## Key Differences from Monitor V1

The main differences between V1 and V2 monitor:

1. **Unified Data Access**: Single ChainStore interface for all blockchain data
2. **Intelligent Caching**: TTL-based strategies reduce RPC load significantly
3. **Capability Awareness**: Graceful handling of different RPC endpoint capabilities
4. **Advanced Analytics**: Real-time metrics with transaction analysis, throughput calculations, and address tracking
5. **Rich UI Experience**: 
   - Dual-pane layout with 8-row atop-style metrics display
   - Interactive navigation with breadcrumb-style page switching
   - Focus-protected modal system for reliable user interactions
   - 10-column enhanced blocks table with sortable columns
6. **Performance Optimized**: String-based method signature matching and efficient counter calculations
7. **Better Architecture**: Clear separation between data access, coordination, analytics, and presentation
8. **4byte Integration**: Automatic transaction signature decoding for better transaction understanding
9. **Future-Ready**: Extensible modal system, metrics framework, and store implementations

The V2 monitor provides significantly richer blockchain analytics, more responsive UI interactions,
and better performance while maintaining efficient RPC usage through intelligent caching. It transforms
basic block monitoring into comprehensive blockchain performance analysis.
