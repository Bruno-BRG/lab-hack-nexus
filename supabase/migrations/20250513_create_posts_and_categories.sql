-- Create posts table
create table if not exists public.posts (
    id uuid default gen_random_uuid() primary key,
    title text not null,
    content text not null,
    slug text not null unique,
    published boolean default false,
    author_id uuid references auth.users(id) on delete cascade,
    category_id uuid references public.categories(id) on delete set null,
    thumbnail_url text,
    created_at timestamp with time zone default timezone('utc'::text, now()) not null,
    updated_at timestamp with time zone default timezone('utc'::text, now()) not null
);

-- Add RLS policies
alter table public.posts enable row level security;

-- Create categories table if it doesn't exist (since posts reference it)
create table if not exists public.categories (
    id uuid default gen_random_uuid() primary key,
    name text not null,
    slug text not null unique,
    description text,
    created_at timestamp with time zone default timezone('utc'::text, now()) not null,
    updated_at timestamp with time zone default timezone('utc'::text, now()) not null
);

-- Add RLS policies for categories
alter table public.categories enable row level security;

-- Policy to allow anyone to read published posts
create policy "Anyone can read published posts"
    on public.posts
    for select
    using (published = true);

-- Policy to allow authenticated users to create posts
create policy "Authenticated users can create posts"
    on public.posts
    for insert
    with check (auth.role() = 'authenticated');

-- Policy to allow post owners to update their posts
create policy "Users can update their own posts"
    on public.posts
    for update
    using (auth.uid() = author_id)
    with check (auth.uid() = author_id);

-- Policy to allow post owners to delete their posts
create policy "Users can delete their own posts"
    on public.posts
    for delete
    using (auth.uid() = author_id);

-- Policy to allow anyone to read categories
create policy "Anyone can read categories"
    on public.categories
    for select
    using (true);

-- Add some sample categories
insert into public.categories (name, slug, description)
values 
    ('Technology', 'technology', 'Posts about technology and innovation'),
    ('Design', 'design', 'Posts about design and creativity'),
    ('Development', 'development', 'Posts about software development')
on conflict (slug) do nothing;
